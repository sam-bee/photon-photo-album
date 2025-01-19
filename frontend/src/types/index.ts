export interface Photo {
  id: string;
  path: string;
  filename: string;
  timestamp: string;
  categories: Category[];
  albums: Album[];
}

export interface Category {
  id: string;
  name: string;
  confidence: number;
}

export interface Album {
  id: string;
  name: string;
  description?: string;
  created_at: string;
  updated_at: string;
}

export interface ApiResponse<T> {
  data: T;
  error?: string;
} 