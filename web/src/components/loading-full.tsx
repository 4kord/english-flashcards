import { Box } from "@mui/system"
import { HashLoader } from "react-spinners"
import { useWindowSize } from "@/hooks/use-window-size";

export const LoadingFull = () => {
    const windowSize = useWindowSize();

    return (
        <Box sx={{
            width: "100%",
            height: windowSize.height,
            display: "flex",
            justifyContent: "center",
            alignItems: "center"
        }}>
            <HashLoader />
        </Box>
    )
}
