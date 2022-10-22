import { Button, Typography } from "@mui/material";
import { Box } from "@mui/system";
import { Link } from "react-router-dom";
import { useAuth } from "../hooks/use-auth"

export const IndexView: React.FC = () => {
  const { auth } = useAuth();

  return (
    <Box>
      <Typography>Index</Typography>
      <Typography>{ auth ? JSON.stringify(auth) : "Not logged in" }</Typography>
      {auth && <Button variant="contained" component={Link} to="/learn">Learn</Button>}
      {auth && auth.role === 1 && <Button variant="contained" component={Link} to="/admin/overview">Admin</Button>}
    </Box>
  )
}
