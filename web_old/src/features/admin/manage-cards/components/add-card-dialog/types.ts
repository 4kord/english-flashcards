export interface IFormData {
  english: string;
  russian: string;
  association: string;
  example: string;
  transcription: string;
  image: FileList;
  audio: FileList;
  imageUrl: string;
  audioUrl: string;
}

export interface IRequest {
  english: string;
  russian: string;
  association: string;
  example: string;
  transcription: string;
  image: File;
  audio: File;
  imageUrl: string;
  audioUrl: string;
}
