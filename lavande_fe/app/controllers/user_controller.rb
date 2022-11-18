class UserController < ApplicationController
  include SessionsHelper
  # before_action :logged_in_user

  def index
    # Check for access token
    token = cookies[:access_token]
    unless token.nil?
      user_data = UserClient.new.get(token)
      unless user_data.nil?
        session[:user_id] = user_data['id']
        @user = User.new(user_data)
      else
        redirect_to new_session_path
      end
    else
      redirect_to new_session_path
    end
  end

  def create
    user_data = params.as_json(:only => [:name, :nik, :email, :phone_no, :religion, :password])
    result = UserClient.new.register(user_data)
    unless result.nil?
      flash[:notice] = "Successfully registered user %s" % params[:name]
      redirect_to new_session_path
    else
      flash[:alert] = "An error occurred when registering user"
    end
  end

  def edit
    # TODO: Prevent request?
    token = cookies[:access_token]
    unless token.nil?
      user_data = UserClient.new.get(token)
      unless user_data.nil?
        session[:user_id] = user_data['id']
        @user = User.new(user_data)
      else
        flash[:alert] = "An error occurred when retrieving user data"
        redirect_to user_index_path
      end
    else
      flash[:alert] = "An error occurred when retrieving user data"
      redirect_to user_index_path
    end
  end

  def update
    token = cookies[:access_token]
    unless token.nil?
      user_data = params[:user].as_json(:only => [:name, :nik, :email, :phone_no, :religion, :password])
      result = UserClient.new.update(user_data, token)
      unless result.nil?
        flash[:notice] = "Successfully updated user data"
        redirect_to user_index_path
      else
        flash[:alert] = "An error occurred when updating user data"
      end
    else
      flash[:alert] = "An error occurred when updating user data"
      redirect_to user_index_path
    end
  end
end
