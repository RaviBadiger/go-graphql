package inmemory

import (
	"errors"
	"fmt"

	"github.com/go-graphql-assignment/graph/model"
)


var products = []model.Product{
	{ID: "1",Title: "Samsung A53",Description: "8GB ", Price: 25000},
	{ID: "2",Title: "iphone 14",Description: "6GB ", Price: 65000},
}

func GetProduct(id string) (*model.Product,error){
	var product model.Product
	var flag bool = false

	for _, item := range products{
		if item.ID == id {
			product = item
			flag = true
			break
		}
	}

	if !flag {
		return nil,errors.New("ID not found")
	}

	return &product, nil
}

func GetProducts()([]*model.Product,error){

	var productCollection []*model.Product

	for index := range products {
		productCollection = append(productCollection, &products[index])
	}	

	return productCollection, nil
}

func CreateProduct(input model.CreateProductInput) (*model.Product,error){
		
		var productItem = model.Product(input)
		
		item, err := GetProduct(productItem.ID)
		fmt.Println(item)
		fmt.Println(err)
		if item != nil {
			return nil, errors.New("product with ID %v" + productItem.ID +" already exists")
		}

		products = append(products, productItem)	
		return &productItem,nil

}

func UpdateProduct(id string, input model.UpdateProductInput)(*model.Product, error){

	var product *model.Product
	for index := range products {
		if products[index].ID == id{
			if input.Title != nil {
				products[index].Title = *input.Title  
			}
			
			if input.Description != nil {
				products[index].Description = *input.Description  
			}
			
			if input.Price != nil {
				products[index].Price = *input.Price  
			}
			product = &products[index]
			return product,nil
		}
	}

	return nil,  errors.New("product with ID %v" + id +" not found")
}

func DeleteProduct(id string)(*model.DeleteProductResponse, error){
	for index := range products {
		if products[index].ID == id{
			products = append(products[:index],products[index+1:]... )
			return &model.DeleteProductResponse{
				DeleteProductID: id,
			},nil
		}
	}
	return &model.DeleteProductResponse{
		DeleteProductID: id,
		}, errors.New("product not found")
}

