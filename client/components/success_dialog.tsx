import ConfirmationDialog from './confirmation_dialog';
import { useState, useEffect } from 'react';

export const SuccessDialog = ({
  firstOpen,
  msg,
  handleClose,
  shouldHistoryBack = true,
}: {
  firstOpen: boolean;
  msg: String;
  handleClose?: () => void;
  shouldHistoryBack?: boolean;
}) => {
  const [open, setOpen] = useState(false);
  useEffect(() => {
    if (firstOpen) {
      setOpen(true);
    }
  }, [firstOpen]);
  const handleSuccessClose = () => {
    setOpen(false);
    if (handleClose) handleClose();
    if (shouldHistoryBack) window.history.back();
  };

  return (
    <ConfirmationDialog
      open={open}
      handleClose={handleSuccessClose}
      msg={msg}
    />
  );
};
