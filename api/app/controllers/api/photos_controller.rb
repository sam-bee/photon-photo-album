module Api
  class PhotosController < ApplicationController
    def index
      photos = Photo.all
      photos = photos.joins(:categories).where(categories: { id: params[:category] }) if params[:category].present?
      photos = photos.joins(:albums).where(albums: { id: params[:album] }) if params[:album].present?
      photos = photos.order(timestamp: params[:sort] == 'date_asc' ? :asc : :desc)

      render json: { data: photos.includes(:categories, :albums) }
    end
  end
end 