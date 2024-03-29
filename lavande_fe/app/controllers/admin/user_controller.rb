class Admin::UserController < ApplicationController
  require_relative '../../clients/admin/user_client'
  PAGINATION_LIMIT = 10

  def index
    @result = {}

    users = Admin::UserClient.new(@token).index(user_list_query)
    if users.success?
      user_data = users.parsed_response["data"]
      users_arr ||= []
      for user in user_data do
        users_arr << User.new(user)
      end
      @result["data"] = users_arr
      @result["meta"] = users.parsed_response["meta"]
      p @result
    else
      err_msg = result.parsed_response["error_message"]
      redirect_back fallback_location: admin_user_index_path, alert: "An error occurred when fetching users: %s" % err_msg
    end
  end

  def status
    payload = params.permit(:user_id, :status)
    payload[:target_id] = payload[:user_id].to_i
    payload.delete(:user_id)

    result = Admin::UserClient.new(@token).status(payload)
    if result.success?
      status_action = "activated"
      if payload["status"] == "nonactive"
        status_action = "deactivated"
      end
      redirect_back fallback_location: admin_user_index_path, notice: "Successfully %s user %s" % [status_action, params[:name]]
    else
      err_msg = result.parsed_response["error_message"]
      redirect_back fallback_location: admin_user_index_path, alert: "An error occurred when updating user status: %s" % err_msg
    end
  end

  def set_admin
    payload = params.permit(:user_id)
    payload[:user_id] = payload[:user_id].to_i
    
    result = Admin::UserClient.new(@token).set_admin(payload)
    if result.success?
      redirect_back fallback_location: admin_user_index_path, notice: "Successfully set user %s as superadmin" % params[:name]
    else
      err_msg = result.parsed_response["error_message"]
      redirect_back fallback_location: admin_user_index_path, alert: "An error occurred when elevating user: %s" % err_msg
    end
  end

  private

  def user_list_query
    params[:page] ||= 1
    params[:page] = params[:page].to_i
    query = params.permit(:page)
    if params[:search].present?
      query[:name] = params[:search]
    end
    query[:limit] = PAGINATION_LIMIT
    query[:offset] = (query[:page] - 1) * PAGINATION_LIMIT
    query
  end
end
