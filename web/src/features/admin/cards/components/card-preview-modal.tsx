import { Button, Dialog, DialogActions, DialogContent, DialogTitle, IconButton, Paper, Typography } from "@mui/material";
import CloseIcon from '@mui/icons-material/Close';
import { Card } from "@/models/card";

interface IProps {
  open: boolean;
  handleClose: () => void;
  card: Card | undefined;
}

export const CardPreviewModal: React.FC<IProps> = ({open, handleClose, card}) => {
  return (
    <Dialog open={open} onClose={handleClose}>
      <DialogTitle sx={{ m: 0, p: 2 }}>
        Card Preview
        <IconButton
          aria-label="close"
          onClick={handleClose}
          sx={{
            position: 'absolute',
            right: 8,
            top: 8,
            color: (theme) => theme.palette.grey[500],
          }}
        >
          <CloseIcon/>
        </IconButton>
      </DialogTitle>
      <DialogContent dividers>
        <Paper variant="outlined" sx={{padding: 2, marginY: 1}}>
          <Typography>English: {card?.english}</Typography>
        </Paper>
        <Paper variant="outlined" sx={{padding: 2, marginY: 1}}>
          <Typography>Russian: {card?.russian}</Typography>
        </Paper>
        { card?.association &&
        <Paper variant="outlined" sx={{padding: 2, marginY: 1}}>
          <Typography>Association: {card.association}</Typography>
        </Paper>
        }
        { card?.example &&
        <Paper variant="outlined" sx={{padding: 2, marginY: 1}}>
          <Typography>Example: {card.example}</Typography>
        </Paper>
        }
        { card?.transcription &&
        <Paper variant="outlined" sx={{padding: 2, marginY: 1}}>
          <Typography>Transcription: {card.transcription}</Typography>
        </Paper>
        }
        { card?.image &&
        <Paper variant="outlined" sx={{padding: 2, marginY: 1}}>
          <img src={card.image_url} alt="previewimg" width="100%"/>
        </Paper>
        }
        { card?.audio &&
        <Paper variant="outlined" sx={{padding: 2, marginY: 1}}>
          <audio controls src={card.audio_url} />
        </Paper>
        }
      </DialogContent>
      <DialogActions>
        <Button onClick={handleClose}>
          Close
        </Button>
      </DialogActions>

    </Dialog>
  );
}
