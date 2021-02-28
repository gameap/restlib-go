#include <iostream>

#include "request.h"
#include "response.h"

#ifdef __cplusplus
extern "C" {
#endif

extern Response DoRequest(Request request);

#ifdef __cplusplus
}
#endif

int main() {
    Request request = {
        "https://ptsv2.com/t/1x32a-1608675136/post",
        Method:POST,
        "body",
        "Accept: application/json\nAuthorization: Bearer token\n",
        "rest-cpp",
        3
    };

    struct Response response = DoRequest(request);

    std::cout << "error: " << response.Error << "\n";
    std::cout << "statusCode: " << response.StatusCode << "\n";
    std::cout << "body: " << response.Body << "\n";
}
