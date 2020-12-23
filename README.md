Simple REST Library writen in GO for use in C/C++ projects

# Usage

```c
#include <stdio.h>
#include <stdlib.h>

#include "request.h"
#include "response.h"

extern Response DoRequest(Request request);

int main() {
    Request request = {
        .Method     = GET,
        .URL        = "http://example.com",
        .Header     = "Accept: application/json",
        .UserAgent  = "rest-c",
    };

    struct Response response = DoRequest(request);

    printf("error: %s \n", response.Error);
    printf("statusCode: %d \n", response.StatusCode);
    printf("body: %s \n", response.Body);
}

```
