import { Box } from "@mui/material";
import { DataGrid, GridColDef } from '@mui/x-data-grid';
import { useCalculateHeight } from "../../../hooks/use-calculate-height";
import { ApiClientMethod, useApi } from "../../../hooks/use-api";
import { IDeck } from "../../../types/deck";
import { useSnackbar } from "notistack";
import { useEffect } from "react";
import { LoadingSmall } from "../../../components/loading-small";
import { useNavigate } from "react-router-dom";
import { NoContent } from "../../../components/nocontent";

const columns: GridColDef[] = [
  { field: 'name', headerName: 'Name', flex: 1},
  { field: 'amount', headerName: 'Amount', type: 'number', flex: 1},
  { field: 'createdAt', headerName: 'Created At', type: "dateTime", flex: 1, valueFormatter: params => new Date(params.value).toUTCString()}
];

export const PremadeDecks: React.FC = () => {
  //
  const navigate = useNavigate();

  // calculated height
  const calculatedHeight = useCalculateHeight()

  // snackbar
  const {enqueueSnackbar} = useSnackbar();

  // fetch decks
  const [fetch, { data, loading }] = useApi<void, Array<IDeck>>({
    method: ApiClientMethod.GET,
    url: "/api/decks/premade",
    onSuccess: (data) => {
      console.log(data);
      enqueueSnackbar("Successfully fetched decks", {variant: "success"});
    },
    onFail: (error) => {
      console.log(error);
      enqueueSnackbar(error.errorCode, {variant: "error"});
    }
  });

  useEffect(() => {
    fetch({});
  }, [fetch])

  return (
    <Box sx={{
      overflowY: "scroll",
    }}>
      {calculatedHeight &&
      <Box sx={{
        padding: 2,
        height: calculatedHeight
      }}>
        <div style={{ height: "100%", width: "100%"}}>
        { data && !loading &&
         <DataGrid
            rows={data}
            columns={columns}
            onRowClick={(data) => navigate(`/premade/${data.id}`)}
            disableSelectionOnClick
          />
        }
        { !data && !loading &&
          <NoContent />
        }
        </div>
      </Box>
      }
      {loading && <LoadingSmall />}
    </Box>
  );
}
