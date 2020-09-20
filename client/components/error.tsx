import Head from 'next/head';
import { NextPageContext } from 'next';
import { CSSProperties } from 'react';

const styles: { [key: string]: CSSProperties } = {
  error: {
    color: '#000',
    background: '#fff',
    fontFamily:
      '-apple-system, BlinkMacSystemFont, Roboto, "Segoe UI", "Fira Sans", Avenir, "Helvetica Neue", "Lucida Grande", sans-serif',
    height: '100vh',
    textAlign: 'center',
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
  },

  container: {
    display: 'flex',
    padding: '0 40px',
    alignItems: 'center',
    justifyContent: 'center',
  },

  desc: {
    display: 'inline-block',
    textAlign: 'left',
    verticalAlign: 'middle',
  },

  h1: {
    display: 'inline-block',
    borderRight: '1px solid rgba(0, 0, 0,.3)',
    margin: 0,
    marginRight: '20px',
    padding: '10px 23px 10px 0',
    fontSize: '24px',
    fontWeight: 500,
    verticalAlign: 'top',
  },

  h2: {
    fontSize: '14px',
    fontWeight: 'normal',
    lineHeight: 'inherit',
    margin: 0,
    padding: 0,
  },
};

export type ErrorProps = {
  statusCode: number;
};

const statusCodeToString = (statusCode: number): string => {
  switch (statusCode) {
    case 401:
      return '認証エラーが発生しました。画面を閉じてからもう一度お試しください。';
    case 404:
      return 'ページが見つかりません。';
    default:
      return '内部エラーが発生しました。';
  }
};

const Error = ({ statusCode }: ErrorProps) => {
  const title = statusCodeToString(statusCode);

  return (
    <div style={styles.error}>
      <Head>
        <title>
          {statusCode}: {title}
        </title>
      </Head>
      <div style={styles.container}>
        <style dangerouslySetInnerHTML={{ __html: 'body { margin: 0 }' }} />
        {statusCode ? <h1 style={styles.h1}>{statusCode}</h1> : null}
        <div style={styles.desc}>
          <h2 style={styles.h2}>{title}</h2>
        </div>
      </div>
    </div>
  );
};

Error.getInitialProps = async ({ res, err }: NextPageContext) => {
  const statusCode =
    res && res.statusCode ? res.statusCode : err ? err.statusCode! : 404;
  return { statusCode };
};

export default Error;
