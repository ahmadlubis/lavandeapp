Rails.application.routes.draw do
  namespace :admin do
    resources :user, only: :index do
      patch 'status'
    end
    resources :unit, only: [:index, :new, :create, :edit, :update]
    resources :tenant, only: [:index, :new, :create] do
      post 'delete', on: :collection
    end
  end

  resources :sessions, only: [:new, :create, :destroy]
  resources :user, only: [:index, :new, :create, :edit, :update]
  resources :unit, only: [:index, :show, :update] do
    resources :tenant, only: [:index, :create] do
      post 'delete', on: :collection
    end
  end

  # patch '/admin/:target_id/status', to: 'admin#status', as: 'admin_status'
  # get '/admin/unit', to: 'admin#unit'
  # get '/admin/unit/new', to: 'admin#unit_new', as: 'admin_new_unit'
  # post '/admin/unit', to: 'admin#unit_create'

  # Define your application routes per the DSL in https://guides.rubyonrails.org/routing.html

  # Defines the root path route ("/")
  # root "articles#index"
  root "user#index"
end
