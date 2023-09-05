package com.rydata.springboot3.repositories.sms;

import com.rydata.springboot3.entities.sms.Sent;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface SentRepository extends JpaRepository<Sent,Integer>{
    
}
