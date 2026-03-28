<?php

namespace App\Controller;

use App\Entity\Product;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\Routing\Annotation\Route;

class ProductController extends AbstractController
{
    #[Route('/products', name: 'product_index', methods: ['GET'])]
    public function index(EntityManagerInterface $em): JsonResponse
    {
        $products = $em->getRepository(Product::class)->findAll();
        $data = [];
        foreach ($products as $product) {
            $data[] = [
                'id' => $product->getId(),
                'name' => $product->getName(),
                'price' => $product->getPrice(),
            ];
        }
        return new JsonResponse($data);
    }

    #[Route('/products', name: 'product_create', methods: ['POST'])]
    public function create(Request $request, EntityManagerInterface $em): JsonResponse
    {
        $data = json_decode($request->getContent(), true);
        $product = new Product();
        $product->setName($data['name']);
        $product->setPrice($data['price']);
        $em->persist($product);
        $em->flush();
        return new JsonResponse(['id' => $product->getId()], 201);
    }

    #[Route('/products/{id}', name: 'product_show', methods: ['GET'])]
    public function show(Product $product): JsonResponse
    {
        return new JsonResponse([
            'id' => $product->getId(),
            'name' => $product->getName(),
            'price' => $product->getPrice(),
        ]);
    }

    #[Route('/products/{id}', name: 'product_update', methods: ['PUT'])]
    public function update(Request $request, Product $product, EntityManagerInterface $em): JsonResponse
    {
        $data = json_decode($request->getContent(), true);
        $product->setName($data['name']);
        $product->setPrice($data['price']);
        $em->flush();
        return new JsonResponse();
    }

    #[Route('/products/{id}', name: 'product_delete', methods: ['DELETE'])]
    public function delete(Product $product, EntityManagerInterface $em): JsonResponse
    {
        $em->remove($product);
        $em->flush();
        return new JsonResponse();
    }
}