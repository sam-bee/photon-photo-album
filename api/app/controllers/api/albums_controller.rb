module Api
  class AlbumsController < ApplicationController
    def index
      render json: { data: Album.all }
    end

    def create
      album = Album.new(album_params)
      if album.save
        render json: { data: album }, status: :created
      else
        render json: { error: album.errors.full_messages }, status: :unprocessable_entity
      end
    end

    def update
      album = Album.find(params[:id])
      if album.update(album_params)
        render json: { data: album }
      else
        render json: { error: album.errors.full_messages }, status: :unprocessable_entity
      end
    end

    def destroy
      Album.find(params[:id]).destroy
      head :no_content
    end

    private

    def album_params
      params.require(:album).permit(:name, :description)
    end
  end
end 