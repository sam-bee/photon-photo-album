import { Photo, Album, ApiResponse } from '../types';

const API_BASE = process.env.REACT_APP_API_URL || 'http://localhost:8080/api';

export const api = {
  async getPhotos(params?: {
    category?: string;
    album?: string;
    sort?: 'date_asc' | 'date_desc';
  }): Promise<ApiResponse<Photo[]>> {
    const queryParams = new URLSearchParams();
    if (params?.category) queryParams.append('category', params.category);
    if (params?.album) queryParams.append('album', params.album);
    if (params?.sort) queryParams.append('sort', params.sort);

    const response = await fetch(`${API_BASE}/photos?${queryParams}`);
    return response.json();
  },

  async getAlbums(): Promise<ApiResponse<Album[]>> {
    const response = await fetch(`${API_BASE}/albums`);
    return response.json();
  },

  async createAlbum(album: Omit<Album, 'id' | 'created_at' | 'updated_at'>): Promise<ApiResponse<Album>> {
    const response = await fetch(`${API_BASE}/albums`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(album),
    });
    return response.json();
  },

  async updateAlbum(id: string, album: Partial<Album>): Promise<ApiResponse<Album>> {
    const response = await fetch(`${API_BASE}/albums/${id}`, {
      method: 'PATCH',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(album),
    });
    return response.json();
  },

  async deleteAlbum(id: string): Promise<ApiResponse<void>> {
    const response = await fetch(`${API_BASE}/albums/${id}`, {
      method: 'DELETE',
    });
    return response.json();
  },
};
