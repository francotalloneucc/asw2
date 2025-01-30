package services

import (
	"context"
	"errors"
	"hotel-api/initializers"
	"hotel-api/models"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Validar si las amenidades existen
func validateAmenitiesExist(amenities []string) error {
	var invalidAmenities []string

	// Recorremos las amenidades y verificamos si existen en la base de datos
	for _, amenity := range amenities {
		var existingAmenity models.Amenity
		err := initializers.DB.Collection("amenities").FindOne(context.Background(), bson.M{"name": amenity}).Decode(&existingAmenity)
		if err != nil {
			invalidAmenities = append(invalidAmenities, amenity)
		}
	}

	if len(invalidAmenities) > 0 {
		// Devolvemos un error si alguna amenidad no existe
		return errors.New("the following amenities do not exist: " + strings.Join(invalidAmenities, ", "))
	}

	return nil
}

// Crear un hotel
func CreateHotel(hotelDto models.Hotel) (models.Hotel, error) {
	// Verificar que las amenidades existan
	if err := validateAmenitiesExist(hotelDto.Amenities); err != nil {
		return models.Hotel{}, err
	}

	hotelDto.ID = primitive.NewObjectID()

	collection := initializers.DB.Collection("hotels")
	_, err := collection.InsertOne(context.Background(), hotelDto)
	if err != nil {
		return models.Hotel{}, err
	}

	return hotelDto, nil
}

// Obtener un hotel por ID
func GetHotel(id primitive.ObjectID) (models.Hotel, error) {
	var hotel models.Hotel
	collection := initializers.DB.Collection("hotels")
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&hotel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Hotel{}, err
		}
		return models.Hotel{}, err
	}

	return hotel, nil
}

// Obtener todos los hoteles
func GetHotels() ([]models.Hotel, error) {
	var hotels []models.Hotel
	collection := initializers.DB.Collection("hotels")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var hotel models.Hotel
		if err := cursor.Decode(&hotel); err != nil {
			return nil, err
		}
		hotels = append(hotels, hotel)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return hotels, nil
}

// Actualizar un hotel
func UpdateHotel(id primitive.ObjectID, hotelDto models.Hotel) (models.Hotel, error) {
	// Verificar que las amenidades existan
	if err := validateAmenitiesExist(hotelDto.Amenities); err != nil {
		return models.Hotel{}, err
	}

	collection := initializers.DB.Collection("hotels")

	update := bson.M{
		"$set": hotelDto,
	}

	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		return models.Hotel{}, err
	}

	return hotelDto, nil
}

// Eliminar un hotel
func DeleteHotel(id primitive.ObjectID) error {
	collection := initializers.DB.Collection("hotels")

	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
