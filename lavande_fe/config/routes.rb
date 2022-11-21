Rails.application.routes.draw do
  get 'admin/index'
  resources :sessions
  resources :user

  # Define your application routes per the DSL in https://guides.rubyonrails.org/routing.html

  # Defines the root path route ("/")
  # root "articles#index"
  root "user#index"
end
