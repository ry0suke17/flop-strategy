import { makeStyles, createStyles } from '@material-ui/core/styles';
import { Fragment } from 'react';
import { CircularProgress, Theme } from '@material-ui/core';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    background: {
      top: 0,
      position: 'absolute',
      width: '100%',
      height: '100%',
      zIndex: theme.zIndex.appBar + 1,
      backgroundColor: '#000000',
      opacity: 0.3,
    },
    loading: {
      position: 'absolute',
      top: '50%',
      left: '50%',
      transform: 'translate(-50%, -50%)',
      zIndex: 1000,
    },
  }),
);

const Loader = ({ open }: { open: boolean }) => {
  const styles = useStyles({});
  return (
    <Fragment>
      {open && (
        <Fragment>
          <div className={styles.background} />
          <div className={styles.loading}>
            <CircularProgress />
          </div>
        </Fragment>
      )}
    </Fragment>
  );
};

export default Loader;
