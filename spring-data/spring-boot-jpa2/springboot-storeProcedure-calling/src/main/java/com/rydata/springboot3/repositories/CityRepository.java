package com.rydata.springboot3.repositories;

import com.rydata.springboot3.entities.City;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface CityRepository extends JpaRepository<City,Long>{

    City findByCityname(String cityname);
    
}
