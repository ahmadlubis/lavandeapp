Rails.application.routes.draw do
  resources :admin, only: :index
  resources :sessions
  resources :user
  resources :unit, only: [:index, :show, :update] do
    resources :tenant, only: [:index, :create]
  end

  patch '/admin/:target_id/status', to: 'admin#status', as: 'admin_status'
  get '/admin/unit', to: 'admin#unit'
  get '/admin/unit/new', to: 'admin#unit_new', as: 'admin_new_unit'
  post '/admin/unit', to: 'admin#unit_create'

  # Define your application routes per the DSL in https://guides.rubyonrails.org/routing.html

  # Defines the root path route ("/")
  # root "articles#index"
  root "user#index"
end
