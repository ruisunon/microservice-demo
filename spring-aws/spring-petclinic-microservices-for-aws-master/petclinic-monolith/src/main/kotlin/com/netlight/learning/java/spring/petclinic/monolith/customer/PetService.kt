/*
 * MIT License
 *
 * Copyright (c) 2020 Alexander Ungar
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package com.netlight.learning.java.spring.petclinic.monolith.customer

import org.springframework.stereotype.Service

@Service
class PetService(
        val petRepository: PetRepository,
        val petTypeRepository: PetTypeRepository
) {
    fun addPetType(type: PetType): PetType = petTypeRepository.save(type)

    fun getPetTypeById(typeId: Long): PetType = petTypeRepository.getOne(typeId)

    fun getAllTypes(): List<PetType> = petTypeRepository.findAll()

    fun addPet(pet: Pet): Pet = petRepository.save(pet)

    fun getPetById(petId: Long): Pet = petRepository.getOne(petId)

    fun updatePet(pet: Pet, petTypeId: Long) {
        val petType = if (petTypeId != pet.type.id) petTypeRepository.getOne(petTypeId) else pet.type
        pet.type = petType
        petRepository.save(pet)
    }

    fun getAllPets(): List<Pet> = petRepository.findAll()
}