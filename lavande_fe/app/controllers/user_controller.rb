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
        session[:role] = user_data['role']
        @user = User.new(user_data)
      else
        redirect_to new_session_path
      end
    else
      redirect_to new_session_path
    end
  end

  def create
    result = UserClient.new.register(user_params)
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
        session[:role] = user_data['role']
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
      p user_params
      result = UserClient.new.update(user_params, token)
      unless result.nil?
        flash[:notice] = "Successfully updated user data"
        redirect_to user_index_path
      else
        flash[:alert] = "An error occurred when updating user data"
        redirect_back(fallback_location: user_index_path)
      end
    else
      flash[:alert] = "An error occurred when updating user data"
      redirect_to user_index_path
    end
  end

  private
 
  def user_params
    params.require(:user).permit(:name, :nik, :email, :phone_no, :role, :status, :religion)
  end

end
