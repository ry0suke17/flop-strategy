import { JSONApiResponse, ModelErrorFromJSON } from '@fls-api-client/src';

export const toModelError = async (error: any) => {
  const errorJSON = new JSONApiResponse(error, (jsonValue) =>
    ModelErrorFromJSON(jsonValue),
  );
  return await errorJSON.value();
};
