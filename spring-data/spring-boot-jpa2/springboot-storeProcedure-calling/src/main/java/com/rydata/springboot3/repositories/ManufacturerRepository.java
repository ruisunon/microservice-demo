package com.rydata.springboot3.repositories;

import com.rydata.springboot3.entities.Manufacturer;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface ManufacturerRepository extends JpaRepository<Manufacturer,Integer>{
    
}
