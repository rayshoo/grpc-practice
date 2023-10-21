import { sendUnaryData, Server, ServerCredentials, ServerUnaryCall } from '@grpc/grpc-js';
import { AuthServiceService } from './protos/auth';
import { LoginCode, LoginRequest, LoginResult } from './protos/auth';

const users = [{ id: 0, username: 'admin', password: 'qwerty' }];

const login = (
  call: ServerUnaryCall<LoginRequest, LoginResult>,
  callback: sendUnaryData<LoginResult>
) => {
  const user = users.find(
    (user) =>
      user.username === call.request.username &&
      user.password === call.request.password
  );

  if (user) {
    const result: LoginResult = {
      loginCode: LoginCode.SUCCESS,
      token: 'RandomSecretToken',
    };
    callback(null, result);
  } else {
    const result: LoginResult = {
      loginCode: LoginCode.FAIL,
    };
    callback(null, result);
  }
};

const server = new Server();
server.addService(AuthServiceService, { login: login });
server.bindAsync('localhost:8080', ServerCredentials.createInsecure(), () => {
  server.start();
});
