#ifndef GDAEMON_RESPONSE_H
#define GDAEMON_RESPONSE_H

typedef struct Response {
    int     StatusCode;
    char*   Status;
    char*   Error;
    char*   Proto;
    char*   Header;
    char*   Body;
    char*   RemoteAddr;
} Response;

#endif //GDAEMON_RESPONSE_H
