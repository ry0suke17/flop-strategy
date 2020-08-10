import Head from 'next/head';
import { theme } from '@fls-lib/material-ui/theme';

const Layout = ({
  children,
  className,
  title = 'This is the default title',
}) => {
  return (
    <div className={className}>
      <Head>
        <title>{title}</title>
        <meta charSet="utf-8" />
        {/* PWA primary color */}
        <meta name="theme-color" content={theme.palette.primary.main} />
        <link
          rel="stylesheet"
          href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,700&display=swap"
        />
        <link
          rel="stylesheet"
          href="https://fonts.googleapis.com/css?family=Armata:300,400,500,700&display=swap"
        />
      </Head>
      {children}
    </div>
  );
};

export default Layout;
