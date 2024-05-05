# frozen_string_literal: true
class Consumer < ApplicationConsumer
  def consume
    messages.each do |message|
      data = message.payload

      existing_product = Product.find_by(id: data["id"])
      if existing_product.nil?
        product = Product.new(
          id: data["id"],
          name: data["name"],
          description: data["description"],
          brand: data["brand"],
          price: data["price"],
          created_at: data["created_at"],
          updated_at: data["updated_at"],
        )

        product.save!
        next
      end

      existing_product.update(data)
      end
    end
  end
end