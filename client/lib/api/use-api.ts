import {
  useState,
  useEffect,
  useReducer,
  Dispatch,
  SetStateAction,
} from 'react';
import { ModelError } from '@fls-api-client/src';
import { toModelError } from './error';

export interface APIResponse<T> {
  response?: T;
  isLoading: boolean;
  error?: ModelError;
}

interface Action<T> {
  type: 'FETCH_INIT' | 'FETCH_SUCCESS' | 'FETCH_FAILURE';
  payload?: T;
  error?: ModelError;
}

const dataFetchReducer = <T>(
  state: APIResponse<T>,
  action: Action<T>,
): APIResponse<T> => {
  switch (action.type) {
    case 'FETCH_INIT':
      return { ...state, isLoading: true, error: undefined };
    case 'FETCH_SUCCESS':
      return {
        ...state,
        isLoading: false,
        response: action.payload,
        error: null,
      };
    case 'FETCH_FAILURE':
      return {
        ...state,
        isLoading: false,
        error: action.error,
      };
    default:
      throw new Error();
  }
};

export const useAPI = <R, T>(
  callFn: (request: R) => Promise<T>,
  initialResponse?: T,
): { state: APIResponse<T>; setRequest: Dispatch<SetStateAction<R>> } => {
  const [request, setRequest] = useState(null);

  const [state, dispatch] = useReducer(dataFetchReducer, {
    isLoading: false,
    response: initialResponse,
  });

  useEffect(() => {
    let didCancel = false;

    const fetchData = async () => {
      if (!request) {
        return;
      }

      dispatch({ type: 'FETCH_INIT' });

      try {
        const result = await callFn(request);

        if (!didCancel) {
          dispatch({ type: 'FETCH_SUCCESS', payload: result });
        }
      } catch (error) {
        if (!didCancel) {
          const modelError = await toModelError(error);
          dispatch({ type: 'FETCH_FAILURE', error: modelError });
        }
      }
    };

    fetchData();

    return () => {
      didCancel = true;
    };
  }, [request]);

  return { state: state as any, setRequest };
};
