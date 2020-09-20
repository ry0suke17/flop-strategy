import App from 'next/app';
import { CssBaseline } from '@material-ui/core';
import { ThemeProvider } from '@material-ui/core/styles';
import { theme } from '@fls-lib/material-ui';
import { Logger } from '@fls-lib/logger';
import { NextPage, NextComponentType, NextPageContext } from 'next';
import Progress from '@fls-components/progress';

Logger.init();

export type PageContext = NextPageContext;

export type Page<P = {}, IP = P> = NextPage<P, IP> & {
  getInitialProps?(context: PageContext): Promise<IP>;
};

class MyApp extends App {
  componentDidMount() {
    // Remove the server-side injected CSS.
    const jssStyles = document.querySelector('#jss-server-side');
    if (jssStyles) {
      jssStyles.parentElement.removeChild(jssStyles);
    }
  }

  render() {
    const { Component, pageProps } = this.props;

    // Workaround for https://github.com/zeit/next.js/issues/8592
    const { err } = this.props as any;
    const modifiedPageProps = { ...pageProps, err };

    return (
      <ThemeProvider theme={theme}>
        {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
        <CssBaseline />
        <Progress />
        <Component {...modifiedPageProps} />
      </ThemeProvider>
    );
  }
}

export default MyApp;
