#include <stdio.h>
#include <stdlib.h>

#include "request.h"
#include "response.h"

extern Response DoRequest(Request request);

int main() {
    Request request = {
        .Method     = POST,
        .URL        = "https://ptsv2.com/t/1x32a-1608675136/post",
        .Body       = "body",
        .Header     = "Accept: application/json",
        .UserAgent  = "rest-c",
        .RetryCount = 3
    };

    struct Response response = DoRequest(request);

    printf("error: %s \n", response.Error);
    printf("statusCode: %d \n", response.StatusCode);
    printf("body: %s \n", response.Body);
}
