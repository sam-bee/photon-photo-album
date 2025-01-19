import React, { useEffect, useState } from 'react';
import styled from '@emotion/styled';
import { Photo, Album, Category } from './types';
import { api } from './api';
import PhotoGrid from './components/PhotoGrid';
import Sidebar from './components/Sidebar';

const AppContainer = styled.div`
  display: flex;
  height: 100vh;
`;

const MainContent = styled.div`
  flex: 1;
  overflow-y: auto;
`;

const App: React.FC = () => {
  const [photos, setPhotos] = useState<Photo[]>([]);
  const [albums, setAlbums] = useState<Album[]>([]);
  const [selectedCategory, setSelectedCategory] = useState<string>();
  const [selectedAlbum, setSelectedAlbum] = useState<string>();

  useEffect(() => {
    const loadPhotos = async () => {
      const response = await api.getPhotos({
        category: selectedCategory,
        album: selectedAlbum,
      });
      if (response.data) {
        setPhotos(response.data);
      }
    };

    const loadAlbums = async () => {
      const response = await api.getAlbums();
      if (response.data) {
        setAlbums(response.data);
      }
    };

    loadPhotos();
    loadAlbums();
  }, [selectedCategory, selectedAlbum]);

  const handleCreateAlbum = async () => {
    const name = prompt('Enter album name:');
    if (name) {
      const response = await api.createAlbum({ name });
      if (response.data) {
        setAlbums([...albums, response.data]);
      }
    }
  };

  // Get unique categories from photos
  const categories = Array.from(
    new Set(photos.flatMap(photo => photo.categories))
  ).sort((a, b) => a.name.localeCompare(b.name));

  return (
    <AppContainer>
      <Sidebar
        categories={categories}
        albums={albums}
        selectedCategory={selectedCategory}
        selectedAlbum={selectedAlbum}
        onCategorySelect={setSelectedCategory}
        onAlbumSelect={setSelectedAlbum}
        onCreateAlbum={handleCreateAlbum}
      />
      <MainContent>
        <PhotoGrid photos={photos} />
      </MainContent>
    </AppContainer>
  );
};

export default App; 