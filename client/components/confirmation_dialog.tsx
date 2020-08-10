import { ReactNode } from 'react';
import {
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
} from '@material-ui/core';

const ConfirmationDialog = ({
  open,
  handleClose,
  title = '確認',
  labelOK = 'OK',
  labelCancel,
  handleOK,
  handleCancel,
  msg,
}: {
  open: boolean;
  handleClose: () => void;
  title?: String;
  labelOK?: String;
  labelCancel?: String;
  handleOK?: () => void;
  handleCancel?: () => void;
  msg: String | ReactNode;
}) => {
  return (
    <Dialog
      fullWidth={true}
      open={open}
      onClose={handleClose}
      aria-labelledby="common-dialog-title"
      aria-describedby="common-dialog-description"
    >
      <DialogTitle id="common-dialog-title">{title}</DialogTitle>
      <DialogContent id="common-dialog-description">{msg}</DialogContent>
      <DialogActions>
        {labelCancel ? (
          <Button onClick={handleCancel || handleClose} color="primary">
            {labelCancel}
          </Button>
        ) : (
          ''
        )}
        <Button onClick={handleOK || handleClose} color="primary">
          {labelOK}
        </Button>
      </DialogActions>
    </Dialog>
  );
};

export default ConfirmationDialog;
