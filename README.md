# Implementation of websocket protocol by golang

## Opening Handshaking docs
 
 - https://datatracker.ietf.org/doc/html/rfc6455#section-4.1

## How are Websocket connections established?
 - https://book.hacktricks.xyz/pentesting-web/cross-site-websocket-hijacking-cswsh#:~:text=Also%20known%20as%20cross%2Dorigin,tokens%20or%20other%20unpredictable%20values.
 - The Connection and Upgrade headers in the request and response indicate that this is a WebSocket handshake.
 - Sec-WebSocket-Accept header in server response:
   - For this header field(Sec-WebSocket-Key), the server has to take the value (as present in the header field, e.g., the base64-encoded [RFC4648] version minus any leading and trailing whitespace) and concatenate this with the Globally Unique Identifier (GUID, [RFC4122]) “258EAFA5-E914-47DA- 95CA-C5AB0DC85B11” in string form, which is unlikely to be used by network endpoints that do not understand the WebSocket Protocol. A SHA-1 hash (160 bits) [FIPS.180-3], base64-encoded (see Section 4 of [RFC4648]), of this concatenation is then returned in the server’s handshake.
 - 

## Websocket Frame Format

       0                   1                   2                   3
      0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
     +-+-+-+-+-------+-+-------------+-------------------------------+
     |F|R|R|R| opcode|M| Payload len |    Extended payload length    |
     |I|S|S|S|  (4)  |A|     (7)     |             (16/64)           |
     |N|V|V|V|       |S|             |   (if payload len==126/127)   |
     | |1|2|3|       |K|             |                               |
     +-+-+-+-+-------+-+-------------+ - - - - - - - - - - - - - - - +
     |     Extended payload length continued, if payload len == 127  |
     + - - - - - - - - - - - - - - - +-------------------------------+
     |                               |Masking-key, if MASK set to 1  |
     +-------------------------------+-------------------------------+
     | Masking-key (continued)       |          Payload Data         |
     +-------------------------------- - - - - - - - - - - - - - - - +
     :                     Payload Data continued ...                :
     + - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - +
     |                     Payload Data continued ...                |
     +---------------------------------------------------------------+   


