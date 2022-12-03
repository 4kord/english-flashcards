export interface AddCard {
  english: string;
  russian: string;
  association?: string;
  example?: string;
  transcription?: string;
  image?: File;
  audio?: File;
  image_url?: string;
  audio_url?: string;
}

export interface Card {
  id: number;
  deck_id: number;
  english: string;
  russian: string;
  association?: string;
  example?: string;
  transcription?: string;
  image?: File;
  audio?: File;
  image_url?: string;
  audio_url?: string;
	created_at?: string;
}
