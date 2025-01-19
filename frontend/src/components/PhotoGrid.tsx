import React from 'react';
import { Photo } from '../types';
import styled from '@emotion/styled';

interface PhotoGridProps {
  photos: Photo[];
  onPhotoClick?: (photo: Photo) => void;
}

const Grid = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 1rem;
  padding: 1rem;
`;

const PhotoItem = styled.div`
  position: relative;
  aspect-ratio: 1;
  cursor: pointer;
  
  &:hover {
    opacity: 0.9;
  }
  
  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    border-radius: 4px;
  }
`;

const PhotoGrid: React.FC<PhotoGridProps> = ({ photos, onPhotoClick }) => {
  return (
    <Grid>
      {photos.map((photo) => (
        <PhotoItem key={photo.id} onClick={() => onPhotoClick?.(photo)}>
          <img src={`/photos/${photo.path}`} alt={photo.filename} loading="lazy" />
        </PhotoItem>
      ))}
    </Grid>
  );
};

export default PhotoGrid; 