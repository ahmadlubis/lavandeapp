class TenantController < ApplicationController
  PAGINATION_LIMIT = 10

  def index
    @result = {}

    unit = UnitClient.new(@token).index(unit_show_query)
    if unit.success?
      @result["unit"] = unit.parsed_response["data"][0]
    else
      err_msg = unit.parsed_response['error_message']
      redirect_back fallback_location: unit_path(params["id"]), alert: "An error occurred when fetching unit: %s" % err_msg
    end

    users = UsersClient.new(@token).index(user_list_query)
    if users.success?
      @result["users"] = users.parsed_response
    else
      err_msg = unit.parsed_response['error_message']
      redirect_back fallback_location: unit_path(params["id"]), alert: "An error occurred when fetching users: %s" % err_msg
    end
  end

  def create
    payload = params.permit(:unit_id, :user_id, :role)
    payload[:unit_id] = payload[:unit_id].to_i
    payload[:user_id] = payload[:user_id].to_i
    payload[:start_at] = Time.now.localtime.iso8601

    result = UnitClient.new(@token).add_tenant(payload)
    if result.success?
      redirect_to unit_path(payload[:unit_id]), notice: "tenant added"
    else
      err_msg = result.parsed_response['error_message']
      redirect_back fallback_location: unit_path(payload[:unit_id]), alert: "An error occurred when adding tenant: %s" % err_msg
    end
  end

  def delete
    payload = params.permit(:unit_id, :user_id)
    payload[:unit_id] = payload[:unit_id].to_i
    payload[:user_id] = payload[:user_id].to_i

    result = UnitClient.new(@token).delete_tenant(payload)
    if result.success?
      redirect_to unit_path(payload[:unit_id]), notice: "tenant deleted"
    else
      err_msg = result.parsed_response['error_message']
      redirect_back fallback_location: unit_path(payload[:unit_id]), alert: "An error occurred when deleting tenant: %s" % err_msg
    end
  end

  private

  def unit_show_query
    query = params.permit(:id)
    query[:limit] = 1
    query[:offset] = 0
    query
  end

  def user_list_query
    params[:page] ||= 1
    params[:page] = params[:page].to_i
    query = params.permit(:page, :name, :email)
    query[:limit] = PAGINATION_LIMIT
    query[:offset] = (query[:page] - 1) * PAGINATION_LIMIT
    query
  end
end
