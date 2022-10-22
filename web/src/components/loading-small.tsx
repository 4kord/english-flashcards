import { CircularProgress } from "@mui/material";
import { Box } from "@mui/system";

export const LoadingSmall: React.FC = () => {
  return (
    <Box sx={{
      width: "100%",
      height: "100%",
      display: "flex",
      justifyContent: "center",
      alignItems: "center",
    }}>
      <CircularProgress />
    </Box>
  );
}
