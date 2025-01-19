import React from 'react';
import styled from '@emotion/styled';
import { Album, Category } from '../types';

interface SidebarProps {
  categories: Category[];
  albums: Album[];
  selectedCategory?: string;
  selectedAlbum?: string;
  onCategorySelect: (category?: string) => void;
  onAlbumSelect: (album?: string) => void;
  onCreateAlbum: () => void;
}

const SidebarContainer = styled.div`
  width: 250px;
  background: #f5f5f5;
  padding: 1rem;
  height: 100vh;
  overflow-y: auto;
`;

const Section = styled.div`
  margin-bottom: 2rem;
  
  h2 {
    font-size: 1.2rem;
    margin-bottom: 1rem;
  }
`;

const Item = styled.div<{ selected?: boolean }>`
  padding: 0.5rem;
  cursor: pointer;
  background: ${props => props.selected ? '#e0e0e0' : 'transparent'};
  border-radius: 4px;
  
  &:hover {
    background: #e0e0e0;
  }
`;

const CreateButton = styled.button`
  width: 100%;
  padding: 0.5rem;
  background: #2196f3;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  
  &:hover {
    background: #1976d2;
  }
`;

const Sidebar: React.FC<SidebarProps> = ({
  categories,
  albums,
  selectedCategory,
  selectedAlbum,
  onCategorySelect,
  onAlbumSelect,
  onCreateAlbum,
}) => {
  return (
    <SidebarContainer>
      <Section>
        <h2>Categories</h2>
        <Item
          selected={!selectedCategory}
          onClick={() => onCategorySelect(undefined)}
        >
          All Photos
        </Item>
        {categories.map((category) => (
          <Item
            key={category.id}
            selected={selectedCategory === category.id}
            onClick={() => onCategorySelect(category.id)}
          >
            {category.name}
          </Item>
        ))}
      </Section>

      <Section>
        <h2>Albums</h2>
        <CreateButton onClick={onCreateAlbum}>Create Album</CreateButton>
        {albums.map((album) => (
          <Item
            key={album.id}
            selected={selectedAlbum === album.id}
            onClick={() => onAlbumSelect(album.id)}
          >
            {album.name}
          </Item>
        ))}
      </Section>
    </SidebarContainer>
  );
};

export default Sidebar; 