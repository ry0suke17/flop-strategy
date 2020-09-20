import React from 'react';
import { default as NextApp, AppProps, AppContext } from 'next/app';
import { Logger } from '@fls-lib/logger';
import ErrorPage from './error';

interface WithErrorProps {
  error?: {
    statusCode: number;
    [key: string]: any;
  };
}

interface AppInitialProps {
  pageProps: WithErrorProps;
}

const withError = function (Error = ErrorPage) {
  return function <P>(WrappedComponent: typeof NextApp) {
    return class WithError extends React.Component<P & AppProps> {
      public static getInitialProps = async (appContext: AppContext) => {
        let appProps: AppInitialProps = { pageProps: {} };

        let error!: WithErrorProps;
        if (WrappedComponent.getInitialProps) {
          try {
            appProps = await WrappedComponent.getInitialProps(appContext);
          } catch (e) {
            await Logger.error(e);

            if (e.error) {
              error = e;
            } else if (e.status) {
              error = {
                error: {
                  statusCode: e.status,
                },
              };
            } else {
              error = {
                error: {
                  statusCode: 500,
                },
              };
            }
            appProps.pageProps = error;
          }
        }

        const { res } = appContext.ctx;
        if (error && res) {
          res.statusCode = error.error?.statusCode ?? 500;
        }
        return appProps;
      };

      render() {
        const { pageProps }: AppInitialProps = this.props;
        const { error } = pageProps;

        if (error && error.statusCode >= 400) {
          const { statusCode, ...additionalErrorProps } = error;
          return (
            <Error statusCode={error.statusCode} {...additionalErrorProps} />
          );
        }

        return <WrappedComponent {...this.props} />;
      }
    };
  };
};

export default withError;
