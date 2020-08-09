import { Configuration, FlopStrategyApi } from '@fls-api-client/src';

// workaround for bugs in the openapi-generator
const globals: any = global;
globals.FormData = require('form-data');
globals.btoa = require('btoa');

const configuration = () => {
  return new Configuration({
    basePath: process.env.BASE_PATH,
    fetchApi: require('isomorphic-unfetch'),
  });
};

export interface API {
  flopStrategy: FlopStrategyApi;
}

export const loadAPI = (): API => {
  const conf = configuration();
  return {
    flopStrategy: new FlopStrategyApi(conf),
  };
};
