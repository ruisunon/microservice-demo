package com.rydata.springboot3.repositories;

import com.rydata.springboot3.entities.Sms;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface SmsRepository extends JpaRepository<Sms,Integer> {
    
}
