import { Backdrop, CircularProgress, Portal } from "@mui/material";
import React from "react";

export const LoadingBackdrop = () => {
  return (
    <Portal>
      <Backdrop
        sx={{ color: '#fff', zIndex: (theme) => theme.zIndex.drawer + 100 }}
        open={true}
      >
        <CircularProgress color="inherit" />
      </Backdrop>
    </Portal>
  );
}
