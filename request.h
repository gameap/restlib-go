#ifndef GDAEMON_REQUEST_H

enum Method {
    GET = 1,
    POST,
    PUT,
    PATCH,
    DELETE,
    HEAD,
    OPTIONS
};

typedef struct Request {
    char*            URL;
    enum Method      Method;
    char*            Body;
    char*            Header;
    char*            UserAgent;
    int              RetryCount;
    int              RetryWaitTime;
} Request;

#define GDAEMON_REQUEST_H

#endif //GDAEMON_REQUEST_H
