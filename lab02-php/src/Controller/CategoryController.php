<?php
namespace App\Controller;

use App\Entity\Category;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\Routing\Annotation\Route;

class CategoryController extends AbstractController {
    #[Route('/categories', methods: ['GET'])]
    public function index(EntityManagerInterface $em): JsonResponse {
        $categories = $em->getRepository(Category::class)->findAll();
        $data = array_map(fn($c) => ['id' => $c->getId(), 'name' => $c->getName(), 'description' => $c->getDescription()], $categories);
        return new JsonResponse($data);
    }

    #[Route('/categories', methods: ['POST'])]
    public function create(Request $request, EntityManagerInterface $em): JsonResponse {
        $data = json_decode($request->getContent(), true);
        $category = new Category();
        $category->setName($data['name']);
        $category->setDescription($data['description'] ?? null);
        $em->persist($category);
        $em->flush();
        return new JsonResponse(['id' => $category->getId()], 201);
    }

    #[Route('/categories/{id}', methods: ['GET'])]
    public function show(Category $category): JsonResponse {
        return new JsonResponse(['id' => $category->getId(), 'name' => $category->getName(), 'description' => $category->getDescription()]);
    }

    #[Route('/categories/{id}', methods: ['PUT'])]
    public function update(Request $request, Category $category, EntityManagerInterface $em): JsonResponse {
        $data = json_decode($request->getContent(), true);
        $category->setName($data['name']);
        $category->setDescription($data['description'] ?? null);
        $em->flush();
        return new JsonResponse();
    }

    #[Route('/categories/{id}', methods: ['DELETE'])]
    public function delete(Category $category, EntityManagerInterface $em): JsonResponse {
        $em->remove($category);
        $em->flush();
        return new JsonResponse();
    }
}
