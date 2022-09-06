// syntax = "proto3";

// package pb;

// import "google/api/annotations.proto";
// import "rpc_create_user.proto";
// import "rpc_login_user.proto";
// import "protoc-gen-openapiv2/options/annotations.proto";

// option go_package = "github.com/rezaAmiri123/simplebank/pb";

// option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = {
//     info: {
//             title: "Simple Bank API";
//             version: "1.1";
//             contact: {
//                     name: "Tech School";
//                     url: "https://github.com/techschool";
//                     email: "techschool.guru@gmail.com";
//             };
//     };
// };


// service SimpleBank {
//     rpc CreateUser (CreateUserRequest) returns (CreateUserResponse) {
//         option (google.api.http) = {
//             post: "/v1/create_user"
//             body: "*"
//         };
//                 option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
//                         description: "Use this API to create a new user";
//                         summary: "Create new user";
//         };
//     }
//     rpc LoginUser (LoginUserRequest) returns (LoginUserResponse) {
//         option (google.api.http) = {
//             post: "/v1/login_user"
//             body: "*"
//         };
//                 option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
//                         description: "Use this API to login user and get access token & refresh token";
//                         summary: "Login user";
//         };
//     }
// }
