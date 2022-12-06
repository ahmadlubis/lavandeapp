class UserController < ApplicationController
  skip_before_action :check_session, only: [:new, :create]

  def index
    result = UserClient.new.get(@token)
    if result.success?
      user_data = result.parsed_response
      session[:user_id] = user_data['id']
      session[:role] = user_data['role']
      @user = User.new(user_data)
    else
      err_msg = result.parsed_response['error_message']
      redirect_to new_session_path, alert: "An error occured: %s" % err_msg
    end
  end

  def new
    unless session[:user_data].nil?
      session[:user_data].as_json().each do |name, value|
        params[name] = value
      end
      session.delete(:user_data)
    end
  end

  def create
    user_data = user_create_params
    session[:user_data] = user_data.except([:password, :password_confirmation])
    # Check password
    if user_data[:password] != user_data[:password_confirmation]
      flash.now[:alert] = "Passwords are not equal"
    else
      result = UserClient.new.register(user_data.except(:password_confirmation))
      if result.success?
        redirect_to new_session_path, notice: "Successfully registered user %s" % user_data[:name]
      else
        err_msg = result.parsed_response['error_message']
        redirect_back fallback_location: new_user_path, alert: "An error occurred when registering user: %s" % err_msg
      end
    end
  end

  def edit
    result = UserClient.new.get(@token)
    if result.success?
      user_data = result.parsed_response
      session[:user_id] = user_data['id']
      session[:role] = user_data['role']
      @user = User.new(user_data)
    else
      redirect_to user_index_path, alert: "An error occurred when retrieving user data"
    end
  end

  def update
    user_data = user_edit_params
    if !user_data[:password].blank?
      # Check password
      if user_data[:password] != user_data[:password_confirmation]
        redirect_back fallback_location: user_index_path, alert: "Passwords are not equal"
        return
      end
    else
      user_data.delete(:password)
    end
    result = UserClient.new.update(user_data.except(:password_confirmation), @token)
    if result.success?
      redirect_to user_index_path, notice: "Successfully updated user data"
    else
      err_msg = result.parsed_response['error_message']
      redirect_back fallback_location: user_index_path, alert: "An error occurred when updating user data: %s" % err_msg
    end
  end

  private
 
  def user_create_params
    params.require([:name, :nik, :email, :password, :password_confirmation, :phone_no, :religion])
    params.permit(:name, :nik, :email, :password, :password_confirmation, :phone_no, :religion)
  end

  def user_edit_params
    params.require(:user).permit(:name, :nik, :email, :password, :password_confirmation, :phone_no, :religion)
  end

end
