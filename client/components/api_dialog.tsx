import { APIResponse } from '@fls-lib/api';
import { Fragment } from 'react';
import Loader from './loader';
import ErrorDialog from './error_dialog';
import { SuccessDialog } from './success_dialog';

const APIDialog = ({
  state,
  msg,
  handleClose,
  shouldHistoryBack,
}: {
  state: APIResponse<any>;
  msg?: String;
  handleClose?: () => void;
  shouldHistoryBack?: boolean;
}) => {
  const { isLoading, error } = state;
  return (
    <Fragment>
      <ErrorDialog error={error} handleClose={handleClose} />
      <Loader open={isLoading} />
      {msg ? (
        <SuccessDialog
          firstOpen={error === null}
          msg={msg}
          handleClose={handleClose}
          shouldHistoryBack={shouldHistoryBack}
        />
      ) : (
        ''
      )}
    </Fragment>
  );
};

export default APIDialog;
