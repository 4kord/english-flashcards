import { Typography } from "@mui/material"
import { Box } from "@mui/system"

export const ForbiddenView: React.FC = () => {

  return (
    <Box>
      <Typography>You can't access this page at the moment (403)</Typography>
    </Box>
  )
}
