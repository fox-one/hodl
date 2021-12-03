// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.4 <0.9.0;

import {BytesLib} from "./bytes.sol";

contract Hodl {
    using BytesLib for bytes;

    event MixinTransaction(bytes);
    event MixinEvent(bytes);

    uint128 public constant PID = 0xd402f2c61ba04d82a75c5674fae1acca;
    uint64 public NONCE = 0;
    mapping(uint128 => uint256) public custodian;
    mapping(address => bytes) public members;

    struct vat {
        uint128 asset; // lock asset
        uint256 amount; // lock amount
        address guy; // the owner
        uint64 end; // vat expiry time [unix epoch time]
        uint64 nonce;
    }

    mapping(uint256 => vat) public vats;
    uint256 tick = 0;

    function lock(
        uint128 asset,
        uint256 amount,
        address guy,
        uint64 end,
        uint64 nonce
    ) internal returns (uint256 id) {
        require(guy != address(0), "guy-not-set");
        require(tick < type(uint256).max, "overflow");

        id = ++tick;
        vats[id].asset = asset;
        vats[id].amount = amount;
        vats[id].guy = guy;
        vats[id].end = end;
        vats[id].nonce = nonce;
    }

    function redeem(
        uint256 id,
        address sender,
        uint64 timestamp
    ) internal {
        require(vats[id].guy == sender, "not-authorized");
        require(vats[id].end < timestamp, "not-end");

        (bytes memory extra, ) = uint256ToVarBytes(id);
        bytes memory log = encodeMixinEvent(
            vats[id].nonce,
            vats[id].asset,
            vats[id].amount,
            extra,
            members[sender]
        );
        emit MixinTransaction(log);

        delete (vats[id]);
    }

    function work(
        address sender,
        uint64 nonce,
        uint128 asset,
        uint256 amount,
        uint64 timestamp,
        bytes memory extra
    ) internal {
        uint256 offset = 0;
        uint256 id = extra.toUint256(offset);
        offset = offset + 32;
        if (id == 0) {
            uint64 exp = extra.toUint64(offset);
            lock(asset, amount, sender, timestamp + exp, nonce);
        } else {
            redeem(id, sender, timestamp);
        }
    }

    // process || nonce || asset || amount || extra || timestamp || members || threshold || sig
    function mixin(bytes calldata raw) public {
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
        bytes memory sig = raw.slice(offset, 64);
        offset = offset + 64;
        require(raw.length == offset, "malformed event encoding");
        // TODO signature verification

        emit MixinEvent(raw);
        custodian[asset] = custodian[asset] + amount;
        members[mixinSenderToAddress(sender)] = sender;

        work(
            mixinSenderToAddress(sender),
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
