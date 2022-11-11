import { Typography } from "@mui/material"
import { Box } from "@mui/system"

export const NotFoundView: React.FC = () => {

  return (
    <Box>
      <Typography>Page not found or file doesn't exist on the server (404)</Typography>
    </Box>
  )
}
