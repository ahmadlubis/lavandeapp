Rails.application.routes.draw do
  resources :admin, only: :index
  resources :sessions
  resources :user

  patch '/admin/:target_id/status', to: 'admin#status', as: 'admin_status'
  get '/admin/unit', to: 'admin#unit'

  # Define your application routes per the DSL in https://guides.rubyonrails.org/routing.html

  # Defines the root path route ("/")
  # root "articles#index"
  root "user#index"
end
