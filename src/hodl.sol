// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0 <0.9.0;

import {BytesLib} from "./bytes.sol";
import {BLS} from "./bls.sol";

contract Hodl {
    using BLS for uint256[2];
    using BLS for bytes;
    using BytesLib for bytes;

    event MixinTransaction(bytes);
    event MixinEvent(
        address indexed sender,
        uint256 nonce,
        uint128 asset,
        uint256 amount,
        uint64 timestamp,
        bytes extra
    );

    uint256[4] public GROUPS = [
        0x2f741961cea2e88cfa2680eeaac040d41f41f3fedb01e38c06f4c6058fd7e425, // x.y
        0x007d68aef83f9690b04f463e13eadd9b18f4869041f1b67e7f1a30c9d1d2c42c, // x.x
        0x2a32fa1736807486256ad8dc6a8740dfb91917cf8d15848133819275be92b673, // y.y
        0x257ad901f02f8a442ccf4f1b1d0d7d3a8e8fe791102706e575d36de1c2a4a40f // y.x
    ];

    uint16 public constant ActionLock = 1;
    uint16 public constant ActionUnlock = 2;

    uint128 public PID = 0xd402f2c61ba04d82a75c5674fae1acca;
    uint64 public NONCE = 0;
    mapping(uint128 => uint256) public custodian;
    mapping(address => bytes) public members;

    struct vat {
        uint128 asset; // lock asset
        uint256 amount; // lock amount
        address guy; // the owner
        uint64 end; // vat expiry time [unix epoch time]
    }

    mapping(uint64 => vat) public vats;

    // --- Init ---
    constructor(uint128 pid_) {
        PID = pid_;
    }

    function lock(
        uint64 id,
        uint128 asset,
        uint256 amount,
        address guy,
        uint64 end
    ) internal {
        require(guy != address(0), "guy-not-set");

        vats[id].asset = asset;
        vats[id].amount = amount;
        vats[id].guy = guy;
        vats[id].end = end;
    }

    function unlock(
        uint64 nonce,
        uint64 id,
        address sender,
        uint64 timestamp,
        bytes memory extra
    ) internal {
        require(vats[id].guy == sender, "not-owner");
        require(vats[id].end < timestamp, "not-end");

        bytes memory log = encodeMixinEvent(
            nonce,
            vats[id].asset,
            vats[id].amount,
            uint128ToFixedBytes(PID).concat(extra),
            members[vats[id].guy]
        );
        emit MixinTransaction(log);
        delete vats[id];
    }

    function decodeExtra(bytes memory extra) internal view returns (uint16 action,uint64 id,uint64 exp) {
        uint256 offset = 0;
        action = extra.toUint16(offset);
        offset += 2;

        id = extra.toUint64(offset);
        offset += 8;

        exp = extra.toUint64(offset);
    }

    function work(
        address sender,
        uint64 nonce,
        uint128 asset,
        uint256 amount,
        uint64 timestamp,
        bytes memory extra
    ) internal returns (bool) {
        (uint16 action,uint64 id,uint64 exp) = decodeExtra(extra);

        if (action == ActionLock) {
            id = nonce + 1;
            lock(id,asset,amount,sender,timestamp+exp);
        } else if (action == ActionUnlock) {
            unlock(nonce,id,sender,timestamp,extra);
        } else {
            return false;
        }

        return true;
    }

    // process || nonce || asset || amount || extra || timestamp || members || threshold || sig
    function mixin(bytes calldata raw) public returns (bool) {
        require(raw.length >= 141, "event data too small");

        uint256 size = 0;
        uint256 offset = 0;
        uint128 process = raw.toUint128(offset);
        require(process == PID, "invalid process");
        offset = offset + 16;

        uint64 nonce = raw.toUint64(offset);
        require(nonce == NONCE, "invalid nonce");
        NONCE = NONCE + 1;
        offset = offset + 8;

        uint128 asset = raw.toUint128(offset);
        offset = offset + 16;

        size = raw.toUint16(offset);
        offset = offset + 2;
        require(size <= 32, "integer out of bounds");
        uint256 amount = new bytes(32 - size)
            .concat(raw.slice(offset, size))
            .toUint256(0);
        offset = offset + size;

        size = raw.toUint16(offset);
        offset = offset + 2;
        bytes memory extra = raw.slice(offset, size);
        offset = offset + size;

        uint64 timestamp = raw.toUint64(offset);
        offset = offset + 8;

        size = raw.toUint16(offset);
        size = 2 + size * 16 + 2;
        bytes memory sender = raw.slice(offset, size);
        offset = offset + size;

        offset = offset + 2;
        require(verifySignature(raw, offset), "invalid signature");

        offset = offset + 64;
        require(raw.length == offset, "malformed event encoding");

        custodian[asset] = custodian[asset] + amount;

        address addr = mixinSenderToAddress(sender);
        members[addr] = sender;

        emit MixinEvent(
            addr,
            nonce,
            asset,
            amount,
            timestamp,
            extra
        );

        return work(
            addr,
            nonce,
            asset,
            amount,
            timestamp,
            extra
        );
    }

    // pid || nonce || asset || amount || extra || timestamp || members || threshold || sig
    function encodeMixinEvent(
        uint64 nonce,
        uint128 asset,
        uint256 amount,
        bytes memory extra,
        bytes memory receiver
    ) internal returns (bytes memory) {
        require(extra.length < 128, "extra too large");
        require(custodian[asset] >= amount, "insufficient custodian");
        custodian[asset] = custodian[asset] - amount;
        bytes memory raw = uint128ToFixedBytes(PID);
        raw = raw.concat(uint64ToFixedBytes(nonce));
        raw = raw.concat(uint128ToFixedBytes(asset));
        (bytes memory ab, uint16 al) = uint256ToVarBytes(amount);
        raw = raw.concat(uint16ToFixedBytes(al));
        raw = raw.concat(ab);
        raw = raw.concat(uint16ToFixedBytes(uint16(extra.length)));
        raw = raw.concat(extra);
        raw = raw.concat(new bytes(8));
        raw = raw.concat(receiver);
        raw = raw.concat(new bytes(2));
        return raw;
    }

    function verifySignature(bytes memory raw, uint256 offset)
        internal
        view
        returns (bool)
    {
        uint256[2] memory sig = [
            raw.toUint256(offset),
            raw.toUint256(offset + 32)
        ];
        uint256[2] memory message = raw
            .slice(0, offset - 2)
            .concat(new bytes(2))
            .hashToPoint();
        return sig.verifySingle(GROUPS, message);
    }

    function mixinSenderToAddress(bytes memory sender)
        internal
        pure
        returns (address)
    {
        return address(uint160(uint256(keccak256(sender))));
    }

    function uint16ToFixedBytes(uint16 x) internal pure returns (bytes memory) {
        bytes memory c = new bytes(2);
        bytes2 b = bytes2(x);
        for (uint256 i = 0; i < 2; i++) {
            c[i] = b[i];
        }
        return c;
    }

    function uint64ToFixedBytes(uint64 x) internal pure returns (bytes memory) {
        bytes memory c = new bytes(8);
        bytes8 b = bytes8(x);
        for (uint256 i = 0; i < 8; i++) {
            c[i] = b[i];
        }
        return c;
    }

    function uint128ToFixedBytes(uint128 x)
        internal
        pure
        returns (bytes memory)
    {
        bytes memory c = new bytes(16);
        bytes16 b = bytes16(x);
        for (uint256 i = 0; i < 16; i++) {
            c[i] = b[i];
        }
        return c;
    }

    function uint256ToVarBytes(uint256 x)
        internal
        pure
        returns (bytes memory, uint16)
    {
        bytes memory c = new bytes(32);
        bytes32 b = bytes32(x);
        uint16 offset = 0;
        for (uint16 i = 0; i < 32; i++) {
            c[i] = b[i];
            if (c[i] > 0 && offset == 0) {
                offset = i;
            }
        }
        uint16 size = 32 - offset;
        return (c.slice(offset, 32 - offset), size);
    }
}
