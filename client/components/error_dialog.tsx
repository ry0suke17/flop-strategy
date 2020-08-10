import ConfirmationDialog from './confirmation_dialog';
import { useState, useEffect } from 'react';
import { ModelError, ModelErrorCodeEnum } from '@fls-api-client/src';
import { Logger } from '@fls-lib/logger';

const implementsModelError = (arg: any): arg is ModelError => {
  return (
    arg !== null &&
    typeof arg === 'object' &&
    (typeof arg.code === 'string' || typeof arg.message === 'string')
  );
};

const errorToString = (error: any): string => {
  if (implementsModelError(error)) {
    switch (error.code) {
      case ModelErrorCodeEnum.CodeInvalidArgument:
        return '無効なパラメータです';
      case undefined:
        return '内部エラーが発生しました';
    }
  }
  return '内部エラーが発生しました';
};

const ErrorDialog = ({
  error,
  handleClose,
}: {
  error: any;
  handleClose?: () => void;
}) => {
  const [errOpen, setErrOpen] = useState(false);
  useEffect(() => {
    if (error) {
      Logger.error(error);
      setErrOpen(true);
    }
  }, [error]);
  const handleErrClose = () => {
    setErrOpen(false);
    if (handleClose) handleClose();
  };
  return (
    <ConfirmationDialog
      open={errOpen}
      title="エラー"
      handleClose={handleErrClose}
      msg={errorToString(error)}
    />
  );
};

export default ErrorDialog;
