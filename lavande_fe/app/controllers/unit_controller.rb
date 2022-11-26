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

  private

  def unit_list_query
    params[:page] ||= 1
    params[:page] = params[:page].to_i
    query = params.permit(:page, :gov_id, :tower, :floor, :unit_no)
    query[:limit] = PAGINATION_LIMIT
    query[:offset] = (query[:page] - 1) * PAGINATION_LIMIT
    query
  end
end
