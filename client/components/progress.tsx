import { useEffect } from 'react';
import NProgress, { NProgressOptions } from 'nprogress';
import Router from 'next/router';

type Props = {
  color: string;
  startPosition: number;
  stopDelayMs: number;
  height: number;
  options?: Partial<NProgressOptions>;
};

const Progress = ({
  color,
  startPosition,
  stopDelayMs,
  height,
  options,
}: Props) => {
  var timer = null;
  useEffect(() => {
    const routeChangeStart = () => {
      NProgress.set(startPosition);
      NProgress.start();
    };
    const routeChangeEnd = () => {
      clearTimeout(timer);
      timer = setTimeout(() => NProgress.done(true), stopDelayMs);
    };

    if (options) NProgress.configure(options);
    Router.events.on('routeChangeStart', routeChangeStart);
    Router.events.on('routeChangeComplete', routeChangeEnd);
    Router.events.on('routeChangeError', () => routeChangeEnd);
  }, []);

  return (
    <style jsx global>{`
      #nprogress {
        pointer-events: none;
      }
      #nprogress .bar {
        background: ${color};
        position: fixed;
        z-index: 1031;
        top: 0;
        left: 0;
        width: 100%;
        height: ${height}px;
      }
      #nprogress .peg {
        display: block;
        position: absolute;
        right: 0px;
        width: 100px;
        height: 100%;
        opacity: 1;
        -webkit-transform: rotate(3deg) translate(0px, -4px);
        -ms-transform: rotate(3deg) translate(0px, -4px);
        transform: rotate(3deg) translate(0px, -4px);
      }
      #nprogress .spinner {
        display: 'block';
        position: fixed;
        z-index: 1031;
        top: 15px;
        right: 15px;
      }
      #nprogress .spinner-icon {
        width: 18px;
        height: 18px;
        box-sizing: border-box;
        border: solid 2px transparent;
        border-top-color: ${color};
        border-left-color: ${color};
        border-radius: 50%;
        -webkit-animation: nprogresss-spinner 400ms linear infinite;
        animation: nprogress-spinner 400ms linear infinite;
      }
      .nprogress-custom-parent {
        overflow: hidden;
        position: relative;
      }
      .nprogress-custom-parent #nprogress .spinner,
      .nprogress-custom-parent #nprogress .bar {
        position: absolute;
      }
      @-webkit-keyframes nprogress-spinner {
        0% {
          -webkit-transform: rotate(0deg);
        }
        100% {
          -webkit-transform: rotate(360deg);
        }
      }
      @keyframes nprogress-spinner {
        0% {
          transform: rotate(0deg);
        }
        100% {
          transform: rotate(360deg);
        }
      }
    `}</style>
  );
};

Progress.defaultProps = {
  color: '#3f51b5',
  startPosition: 0.3,
  stopDelayMs: 100,
  height: 2,
  options: {
    showSpinner: false,
  },
};

export default Progress;
