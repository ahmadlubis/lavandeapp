class UnitController < ApplicationController
  PAGINATION_LIMIT = 10

  def index
    result = UnitClient.new(@token).index(unit_list_query)
    if result.success?
      @result = result.parsed_response
    else
      @err_msg = result.parsed_response['error_message']
    end
  end

  def show
    @result = {}
    client = UnitClient.new(@token)

    unit = client.index(unit_show_query)
    if unit.success?
      @result["unit"] = unit.parsed_response["data"][0]
    else
      err_msg = unit.parsed_response['error_message']
      redirect_back fallback_location: unit_index_path, alert: "An error occurred when fetching unit: %s" % err_msg
      return
    end

    tenants = client.index_tenant(tenant_list_query)
    if tenants.success?
      @result["tenants"] = tenants.parsed_response
    else
      err_msg = tenants.parsed_response['error_message']
      redirect_back fallback_location: unit_index_path, alert: "An error occurred when fetching tenants: %s" % err_msg
    end

    @result["tenants"]["data"].each do |v|
      user = UsersClient.new(@token).index({id: v["user_id"], limit: 1, offset: 0})
      if user.success?
        v["user"] = user.parsed_response["data"][0]
      else
        err_msg = tenants.parsed_response['error_message']
        redirect_back fallback_location: unit_index_path, alert: "An error occurred when fetching tenants: %s" % err_msg
      end
    end
  end

  def update
    payload = params.permit(:id, :ajb, :akte)
    payload[:id] = payload[:id].to_i
    if params[:ajb].present?
      if params[:ajb].content_type != "application/pdf"
        redirect_back fallback_location: unit_path(params[:id]), alert: "AJB should be a PDF File"
        return
      end
      payload[:ajb] = Base64.encode64(params[:ajb].read)
    end
    if params[:akte].present?
      if params[:akte].content_type != "application/pdf"
        redirect_back fallback_location: unit_path(params[:id]), alert: "Akte should be a PDF File"
        return
      end
      payload[:akte] = Base64.encode64(params[:akte].read)
    end

    result = UnitClient.new(@token).update(payload)
    if result.success?
      redirect_back fallback_location: unit_path(params[:id]), notice: "unit updated"
    else
      err_msg = result.parsed_response['error_message']
      redirect_back fallback_location: unit_path(params[:id]), alert: "An error occurred when updating unit: %s" % err_msg
    end
  end

  private

  def unit_list_query
    params[:page] ||= 1
    params[:page] = params[:page].to_i
    query = params.permit(:page, :gov_id, :tower, :floor, :unit_no)
    query[:limit] = PAGINATION_LIMIT
    query[:offset] = (query[:page] - 1) * PAGINATION_LIMIT
    query
  end

  def unit_show_query
    query = params.permit(:id)
    query[:limit] = 1
    query[:offset] = 0
    query
  end

  def tenant_list_query
    params[:page] ||= 1
    params[:page] = params[:page].to_i
    query = params.permit(:id, :page)
    query[:unit_id] = query[:id].to_i
    query[:limit] = PAGINATION_LIMIT
    query[:offset] = (query[:page] - 1) * PAGINATION_LIMIT
    query
  end
end
